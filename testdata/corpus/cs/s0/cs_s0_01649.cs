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

    public interface IHandler15430
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15430(string Label, int Payload);

    public sealed class Handler15430 : IHandler15430
    {
        public void Handle(int value) { }
    }
    public interface IHandler15431
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15431(string Label, int Payload);

    public sealed class Handler15431 : IHandler15431
    {
        public void Handle(int value) { }
    }
    public static int Guard15432(int value)
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
    public interface IHandler15433
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15433(string Label, int Payload);

    public sealed class Handler15433 : IHandler15433
    {
        public void Handle(int value) { }
    }
    public struct Point15434
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 13;
    }

    public static class Grid15434
    {
        public static Point15434 Make(int left, int right)
        {
            return new Point15434 { X = left, Y = right };
        }
    }
}
