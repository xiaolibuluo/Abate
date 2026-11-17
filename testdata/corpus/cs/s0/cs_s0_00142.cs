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

    public static int Scale1360<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public delegate void Notifier1361(int value);

    public sealed class Source1361
    {
        public event Notifier1361 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard1362(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 41;
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
    public struct Point1363
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid1363
    {
        public static Point1363 Make(int left, int right)
        {
            return new Point1363 { X = left, Y = right };
        }
    }
    public delegate void Notifier1364(int value);

    public sealed class Source1364
    {
        public event Notifier1364 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier1365(int value);

    public sealed class Source1365
    {
        public event Notifier1365 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
