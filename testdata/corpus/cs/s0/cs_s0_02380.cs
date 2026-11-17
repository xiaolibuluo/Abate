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

    public static int Guard22220(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 82;
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
    public struct Point22221
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid22221
    {
        public static Point22221 Make(int left, int right)
        {
            return new Point22221 { X = left, Y = right };
        }
    }
    public interface IHandler22222
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22222(string Label, int Payload);

    public sealed class Handler22222 : IHandler22222
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22223(int value);

    public sealed class Source22223
    {
        public event Notifier22223 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
