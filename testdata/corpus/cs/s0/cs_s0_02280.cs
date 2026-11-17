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

    public static int Guard21290(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 37;
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
    public delegate void Notifier21291(int value);

    public sealed class Source21291
    {
        public event Notifier21291 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21292<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public struct Point21293
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 58;
    }

    public static class Grid21293
    {
        public static Point21293 Make(int left, int right)
        {
            return new Point21293 { X = left, Y = right };
        }
    }
    public static int Scale21294<T>(T value) where T : struct
    {
        int number = 65;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
}
