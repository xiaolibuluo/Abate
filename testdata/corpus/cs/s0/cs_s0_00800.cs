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

    public interface IHandler7510
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7510(string Label, int Payload);

    public sealed class Handler7510 : IHandler7510
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate7511(IEnumerable<int> source)
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
    public struct Point7512
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 82;
    }

    public static class Grid7512
    {
        public static Point7512 Make(int left, int right)
        {
            return new Point7512 { X = left, Y = right };
        }
    }
}
