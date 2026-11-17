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

    public interface IHandler12840
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12840(string Label, int Payload);

    public sealed class Handler12840 : IHandler12840
    {
        public void Handle(int value) { }
    }
    public static int Guard12841(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 20;
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
    public struct Point12842
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid12842
    {
        public static Point12842 Make(int left, int right)
        {
            return new Point12842 { X = left, Y = right };
        }
    }
}
