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

    public static int Guard17240(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 10;
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
    public interface IHandler17241
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17241(string Label, int Payload);

    public sealed class Handler17241 : IHandler17241
    {
        public void Handle(int value) { }
    }
    public struct Point17242
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 24;
    }

    public static class Grid17242
    {
        public static Point17242 Make(int left, int right)
        {
            return new Point17242 { X = left, Y = right };
        }
    }
}
