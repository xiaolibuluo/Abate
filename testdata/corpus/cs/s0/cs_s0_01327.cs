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

    public struct Point12430
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 62;
    }

    public static class Grid12430
    {
        public static Point12430 Make(int left, int right)
        {
            return new Point12430 { X = left, Y = right };
        }
    }
    public interface IHandler12431
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12431(string Label, int Payload);

    public sealed class Handler12431 : IHandler12431
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate12432(IEnumerable<int> source)
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
