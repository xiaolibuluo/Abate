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

    public delegate void Notifier4300(int value);

    public sealed class Source4300
    {
        public event Notifier4300 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard4301(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 73;
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
    public interface IHandler4302
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4302(string Label, int Payload);

    public sealed class Handler4302 : IHandler4302
    {
        public void Handle(int value) { }
    }
    public interface IHandler4303
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4303(string Label, int Payload);

    public sealed class Handler4303 : IHandler4303
    {
        public void Handle(int value) { }
    }
    public interface IHandler4304
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4304(string Label, int Payload);

    public sealed class Handler4304 : IHandler4304
    {
        public void Handle(int value) { }
    }
    public struct Point4305
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid4305
    {
        public static Point4305 Make(int left, int right)
        {
            return new Point4305 { X = left, Y = right };
        }
    }
}
