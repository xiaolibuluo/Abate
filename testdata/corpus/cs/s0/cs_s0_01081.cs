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

    public async Task<int> Aggregate10130(IEnumerable<int> source)
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
    public struct Point10131
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid10131
    {
        public static Point10131 Make(int left, int right)
        {
            return new Point10131 { X = left, Y = right };
        }
    }
    public static int Scale10132<T>(T value) where T : struct
    {
        int number = 93;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public interface IHandler10133
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10133(string Label, int Payload);

    public sealed class Handler10133 : IHandler10133
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10134(IEnumerable<int> source)
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
