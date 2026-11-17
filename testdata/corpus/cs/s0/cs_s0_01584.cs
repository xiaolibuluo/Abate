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

    public static int Guard14820(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 65;
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
    public interface IHandler14821
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14821(string Label, int Payload);

    public sealed class Handler14821 : IHandler14821
    {
        public void Handle(int value) { }
    }
    public struct Point14822
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 79;
    }

    public static class Grid14822
    {
        public static Point14822 Make(int left, int right)
        {
            return new Point14822 { X = left, Y = right };
        }
    }
}
