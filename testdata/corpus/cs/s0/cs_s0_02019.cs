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

    public static int Guard18830(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 40;
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
    public async Task<int> Aggregate18831(IEnumerable<int> source)
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
    public interface IHandler18832
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18832(string Label, int Payload);

    public sealed class Handler18832 : IHandler18832
    {
        public void Handle(int value) { }
    }
    public struct Point18833
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid18833
    {
        public static Point18833 Make(int left, int right)
        {
            return new Point18833 { X = left, Y = right };
        }
    }
}
