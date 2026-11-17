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

    public static int Scale10670<T>(T value) where T : struct
    {
        int number = 5;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public interface IHandler10671
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10671(string Label, int Payload);

    public sealed class Handler10671 : IHandler10671
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10672(IEnumerable<int> source)
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
    public struct Point10673
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid10673
    {
        public static Point10673 Make(int left, int right)
        {
            return new Point10673 { X = left, Y = right };
        }
    }
    public struct Point10674
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid10674
    {
        public static Point10674 Make(int left, int right)
        {
            return new Point10674 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate10675(IEnumerable<int> source)
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
