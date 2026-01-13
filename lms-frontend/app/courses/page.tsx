"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { apiFetch } from "@/lib/api";

type Course = {
  id: number;
  title: string;
  description: string;
};

export default function CoursesPage() {
  const [courses, setCourses] = useState<Course[]>([]);

  useEffect(() => {
    apiFetch("/courses")
      .then(res => res.json())
      .then(setCourses)
      .catch(console.error);
  }, []);

  return (
    <main className="min-h-screen bg-black text-white p-6">
      <h1 className="text-3xl font-bold mb-6">Daftar Course</h1>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {courses.map(course => (
          <Link
            key={course.id}
            href={`/courses/${course.id}`}
            className="block rounded-xl border border-white/20 p-4
                       hover:bg-white/10 hover:border-white/40
                       transition"
          >
            <h2 className="text-xl font-semibold mb-2">
              {course.title}
            </h2>
            <p className="text-sm text-white/70">
              {course.description}
            </p>
          </Link>
        ))}
      </div>
    </main>
  );
}
