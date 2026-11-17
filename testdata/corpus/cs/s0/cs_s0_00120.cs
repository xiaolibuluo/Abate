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

    public delegate void Notifier1170(int value);

    public sealed class Source1170
    {
        public event Notifier1170 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point1171
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid1171
    {
        public static Point1171 Make(int left, int right)
        {
            return new Point1171 { X = left, Y = right };
        }
    }
    public static int Guard1172(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 85;
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
