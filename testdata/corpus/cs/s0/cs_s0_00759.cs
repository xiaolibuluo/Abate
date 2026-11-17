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

    public delegate void Notifier7110(int value);

    public sealed class Source7110
    {
        public event Notifier7110 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier7111(int value);

    public sealed class Source7111
    {
        public event Notifier7111 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard7112(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 47;
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
    public struct Point7113
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 54;
    }

    public static class Grid7113
    {
        public static Point7113 Make(int left, int right)
        {
            return new Point7113 { X = left, Y = right };
        }
    }
}
