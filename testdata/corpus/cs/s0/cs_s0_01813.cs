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

    public struct Point16920
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 79;
    }

    public static class Grid16920
    {
        public static Point16920 Make(int left, int right)
        {
            return new Point16920 { X = left, Y = right };
        }
    }
    public delegate void Notifier16921(int value);

    public sealed class Source16921
    {
        public event Notifier16921 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard16922(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 93;
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
    public struct Point16923
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid16923
    {
        public static Point16923 Make(int left, int right)
        {
            return new Point16923 { X = left, Y = right };
        }
    }
    public interface IHandler16924
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16924(string Label, int Payload);

    public sealed class Handler16924 : IHandler16924
    {
        public void Handle(int value) { }
    }
}
