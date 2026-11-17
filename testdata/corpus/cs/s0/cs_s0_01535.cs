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

    public interface IHandler14390
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14390(string Label, int Payload);

    public sealed class Handler14390 : IHandler14390
    {
        public void Handle(int value) { }
    }
    public static int Guard14391(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 95;
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
    public interface IHandler14392
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14392(string Label, int Payload);

    public sealed class Handler14392 : IHandler14392
    {
        public void Handle(int value) { }
    }
    public struct Point14393
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 12;
    }

    public static class Grid14393
    {
        public static Point14393 Make(int left, int right)
        {
            return new Point14393 { X = left, Y = right };
        }
    }
}
