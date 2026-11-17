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

    public delegate void Notifier26200(int value);

    public sealed class Source26200
    {
        public event Notifier26200 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier26201(int value);

    public sealed class Source26201
    {
        public event Notifier26201 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point26202
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 32;
    }

    public static class Grid26202
    {
        public static Point26202 Make(int left, int right)
        {
            return new Point26202 { X = left, Y = right };
        }
    }
    public delegate void Notifier26203(int value);

    public sealed class Source26203
    {
        public event Notifier26203 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard26204(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 46;
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
    public delegate void Notifier26205(int value);

    public sealed class Source26205
    {
        public event Notifier26205 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
