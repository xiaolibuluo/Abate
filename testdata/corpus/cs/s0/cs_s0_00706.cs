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

    public struct Point6610
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 62;
    }

    public static class Grid6610
    {
        public static Point6610 Make(int left, int right)
        {
            return new Point6610 { X = left, Y = right };
        }
    }
    public static int Guard6611(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 69;
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
    public interface IHandler6612
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6612(string Label, int Payload);

    public sealed class Handler6612 : IHandler6612
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate6613(IEnumerable<int> source)
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
}
