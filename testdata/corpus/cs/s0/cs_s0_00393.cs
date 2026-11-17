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

    public static int Guard3720(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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
    public delegate void Notifier3721(int value);

    public sealed class Source3721
    {
        public event Notifier3721 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler3722
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3722(string Label, int Payload);

    public sealed class Handler3722 : IHandler3722
    {
        public void Handle(int value) { }
    }
    public struct Point3723
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 12;
    }

    public static class Grid3723
    {
        public static Point3723 Make(int left, int right)
        {
            return new Point3723 { X = left, Y = right };
        }
    }
}
