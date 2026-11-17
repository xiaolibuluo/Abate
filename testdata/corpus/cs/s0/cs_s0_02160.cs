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

    public async Task<int> Aggregate20160(IEnumerable<int> source)
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
    public interface IHandler20161
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20161(string Label, int Payload);

    public sealed class Handler20161 : IHandler20161
    {
        public void Handle(int value) { }
    }
    public struct Point20162
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid20162
    {
        public static Point20162 Make(int left, int right)
        {
            return new Point20162 { X = left, Y = right };
        }
    }
}
