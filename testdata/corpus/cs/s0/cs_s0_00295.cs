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

    public interface IHandler2780
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2780(string Label, int Payload);

    public sealed class Handler2780 : IHandler2780
    {
        public void Handle(int value) { }
    }
    public struct Point2781
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid2781
    {
        public static Point2781 Make(int left, int right)
        {
            return new Point2781 { X = left, Y = right };
        }
    }
    public struct Point2782
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid2782
    {
        public static Point2782 Make(int left, int right)
        {
            return new Point2782 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate2783(IEnumerable<int> source)
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
