import * as React from 'react'
import { Link } from 'react-router-dom'
import Wrapper from '../components/shared/PageWrapper'
import PageHeader from '../components/shared/PageHeader'

const About = () => (
  <Wrapper>
    <PageHeader />
    <Link to='/'>Go to home page</Link>
  </Wrapper>
)

export default About
