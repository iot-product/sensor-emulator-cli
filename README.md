# sensor-emulator-cli
## Motivation
At work, I usually work wih some sensor reading and shows them at a dashboard created by 
frontend team (I'm a IoT backend engineer to add some context). So, instead of using some sensor for testing, 
I create this emulator to emulate the behaviour on ideal condition.
## Specification
<table>
  <tr>
    <th>Tools</th>
    <th>Description</th>
  <tr>
    <td>Broker</td>
    <td>broker.hivemq.com</td>
  </tr>
  <tr>
    <td>Port</td>
    <td>1883</td>
  </tr>
 </tr>
</table>

## Documentation
### LM35
LM 35 is a temperature sensor. Based on this [datasheet](https://www.ti.com/lit/ds/symlink/lm35.pdf?HQS=TI-null-null-alldatasheets-df-pf-SEP-wwe)
. It has reading range of -55 - 150 degrees celcius.
<table>
  <tr>
    <th>Flags</th>
    <th>Description</th>
  <tr>
    <td>interval</td>
    <td>integer, set interval of data transmitted</td>
  </tr>
  <tr>
    <td>Topic</td>
    <td>Input topic endpoint for publishing</td>
  </tr>
 </tr>
</table>

Example:
```bash
# publish data to topic/test every 2 seconds
sensim lm35 --interval=2 --topic=topic/test
```