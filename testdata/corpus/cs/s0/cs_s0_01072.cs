using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public struct Point10040
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid10040
    {
        public static Point10040 Make(int left, int right)
        {
            return new Point10040 { X = left, Y = right };
        }
    }
    public delegate void Notifier10041(int value);

    public sealed class Source10041
    {
        public event Notifier10041 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point10042
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 73;
    }

    public static class Grid10042
    {
        public static Point10042 Make(int left, int right)
        {
            return new Point10042 { X = left, Y = right };
        }
    }
    public static int Guard10043(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 80;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
}
