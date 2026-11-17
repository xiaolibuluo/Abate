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

    public struct Point19400
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid19400
    {
        public static Point19400 Make(int left, int right)
        {
            return new Point19400 { X = left, Y = right };
        }
    }
    public static int Guard19401(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 12;
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
    public delegate void Notifier19402(int value);

    public sealed class Source19402
    {
        public event Notifier19402 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19403(int value);

    public sealed class Source19403
    {
        public event Notifier19403 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19404(int value);

    public sealed class Source19404
    {
        public event Notifier19404 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point19405
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid19405
    {
        public static Point19405 Make(int left, int right)
        {
            return new Point19405 { X = left, Y = right };
        }
    }
}
