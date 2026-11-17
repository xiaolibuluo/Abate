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

    public static int Guard17260(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 36;
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
    public interface IHandler17261
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17261(string Label, int Payload);

    public sealed class Handler17261 : IHandler17261
    {
        public void Handle(int value) { }
    }
    public struct Point17262
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid17262
    {
        public static Point17262 Make(int left, int right)
        {
            return new Point17262 { X = left, Y = right };
        }
    }
}
