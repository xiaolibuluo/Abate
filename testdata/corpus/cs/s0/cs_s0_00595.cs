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

    public interface IHandler5560
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5560(string Label, int Payload);

    public sealed class Handler5560 : IHandler5560
    {
        public void Handle(int value) { }
    }
    public struct Point5561
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 62;
    }

    public static class Grid5561
    {
        public static Point5561 Make(int left, int right)
        {
            return new Point5561 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate5562(IEnumerable<int> source)
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
