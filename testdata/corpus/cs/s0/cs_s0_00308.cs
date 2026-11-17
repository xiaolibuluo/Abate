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

    public async Task<int> Aggregate2910(IEnumerable<int> source)
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
    public async Task<int> Aggregate2911(IEnumerable<int> source)
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
    public delegate void Notifier2912(int value);

    public sealed class Source2912
    {
        public event Notifier2912 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler2913
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2913(string Label, int Payload);

    public sealed class Handler2913 : IHandler2913
    {
        public void Handle(int value) { }
    }
    public struct Point2914
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid2914
    {
        public static Point2914 Make(int left, int right)
        {
            return new Point2914 { X = left, Y = right };
        }
    }
    public struct Point2915
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid2915
    {
        public static Point2915 Make(int left, int right)
        {
            return new Point2915 { X = left, Y = right };
        }
    }
}
