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

    public delegate void Notifier23130(int value);

    public sealed class Source23130
    {
        public event Notifier23130 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23131
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid23131
    {
        public static Point23131 Make(int left, int right)
        {
            return new Point23131 { X = left, Y = right };
        }
    }
    public interface IHandler23132
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23132(string Label, int Payload);

    public sealed class Handler23132 : IHandler23132
    {
        public void Handle(int value) { }
    }
    public static int Guard23133(int value)
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
}
