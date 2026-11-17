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

    public struct Point390
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid390
    {
        public static Point390 Make(int left, int right)
        {
            return new Point390 { X = left, Y = right };
        }
    }
    public struct Point391
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid391
    {
        public static Point391 Make(int left, int right)
        {
            return new Point391 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate392(IEnumerable<int> source)
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
    public static int Scale393<T>(T value) where T : struct
    {
        int number = 48;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public interface IHandler394
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope394(string Label, int Payload);

    public sealed class Handler394 : IHandler394
    {
        public void Handle(int value) { }
    }
}
