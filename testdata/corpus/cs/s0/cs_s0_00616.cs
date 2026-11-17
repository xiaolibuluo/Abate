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

    public delegate void Notifier5750(int value);

    public sealed class Source5750
    {
        public event Notifier5750 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5751<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public static int Guard5752(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
    public delegate void Notifier5753(int value);

    public sealed class Source5753
    {
        public event Notifier5753 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5754<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public struct Point5755
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 46;
    }

    public static class Grid5755
    {
        public static Point5755 Make(int left, int right)
        {
            return new Point5755 { X = left, Y = right };
        }
    }
}
