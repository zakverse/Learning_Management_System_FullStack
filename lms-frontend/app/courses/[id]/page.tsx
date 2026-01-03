"use client";

import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { apiFetch } from "@/lib/api";

export default function CourseDetail() {
  const { id } = useParams();
  const [course, setCourse] = useState<any>(null);

  useEffect(() => {
    apiFetch(`/courses/${id}`)
      .then(res => res.json())
      .then(setCourse);
  }, [id]);

  if (!course) return <p className="text-white p-6">Loading...</p>;

  return (
    <main className="p-6 text-white">
      <h1 className="text-2xl font-bold">{course.title}</h1>
      <p className="mt-2">{course.description}</p>
    </main>
  );
}
