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

    public interface IHandler8100
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8100(string Label, int Payload);

    public sealed class Handler8100 : IHandler8100
    {
        public void Handle(int value) { }
    }
    public struct Point8101
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 66;
    }

    public static class Grid8101
    {
        public static Point8101 Make(int left, int right)
        {
            return new Point8101 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate8102(IEnumerable<int> source)
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
    public static int Guard8103(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 80;
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
