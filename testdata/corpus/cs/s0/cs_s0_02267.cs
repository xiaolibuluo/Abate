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

    public struct Point21170
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 75;
    }

    public static class Grid21170
    {
        public static Point21170 Make(int left, int right)
        {
            return new Point21170 { X = left, Y = right };
        }
    }
    public delegate void Notifier21171(int value);

    public sealed class Source21171
    {
        public event Notifier21171 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21172
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21172(string Label, int Payload);

    public sealed class Handler21172 : IHandler21172
    {
        public void Handle(int value) { }
    }
    public static int Guard21173(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 96;
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
