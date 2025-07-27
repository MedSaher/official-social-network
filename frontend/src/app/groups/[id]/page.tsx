// app/groups/[id]/page.tsx

import React from 'react';
import JoinRequestsClient from './JoinRequestsClient';

interface GroupPageProps {
  params: { id: string };
}

export default function GroupPage({ params }: GroupPageProps) {
  const { id } = params;

  return (
    <div>
      <h1>Group ID: {id}</h1>
      <h2>Pending Join Requests</h2>
      {/* Pass group id to client component */}
      <JoinRequestsClient groupId={id} />
    </div>
  );
}
