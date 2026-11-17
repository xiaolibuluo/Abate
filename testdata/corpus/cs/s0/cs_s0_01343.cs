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

    public struct Point12580
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid12580
    {
        public static Point12580 Make(int left, int right)
        {
            return new Point12580 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate12581(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public interface IHandler12582
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12582(string Label, int Payload);

    public sealed class Handler12582 : IHandler12582
    {
        public void Handle(int value) { }
    }
    public interface IHandler12583
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12583(string Label, int Payload);

    public sealed class Handler12583 : IHandler12583
    {
        public void Handle(int value) { }
    }
    public static int Guard12584(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 91;
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
