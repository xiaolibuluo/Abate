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

    public interface IHandler12420
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12420(string Label, int Payload);

    public sealed class Handler12420 : IHandler12420
    {
        public void Handle(int value) { }
    }
    public struct Point12421
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 56;
    }

    public static class Grid12421
    {
        public static Point12421 Make(int left, int right)
        {
            return new Point12421 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate12422(IEnumerable<int> source)
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
