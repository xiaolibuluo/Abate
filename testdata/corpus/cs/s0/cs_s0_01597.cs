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

    public static int Scale14930<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public async Task<int> Aggregate14931(IEnumerable<int> source)
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
    public interface IHandler14932
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14932(string Label, int Payload);

    public sealed class Handler14932 : IHandler14932
    {
        public void Handle(int value) { }
    }
    public struct Point14933
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 35;
    }

    public static class Grid14933
    {
        public static Point14933 Make(int left, int right)
        {
            return new Point14933 { X = left, Y = right };
        }
    }
}
