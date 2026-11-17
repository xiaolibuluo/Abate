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

    public interface IHandler12130
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12130(string Label, int Payload);

    public sealed class Handler12130 : IHandler12130
    {
        public void Handle(int value) { }
    }
    public static int Guard12131(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 67;
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
    public struct Point12132
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid12132
    {
        public static Point12132 Make(int left, int right)
        {
            return new Point12132 { X = left, Y = right };
        }
    }
    public struct Point12133
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid12133
    {
        public static Point12133 Make(int left, int right)
        {
            return new Point12133 { X = left, Y = right };
        }
    }
    public static int Guard12134(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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
