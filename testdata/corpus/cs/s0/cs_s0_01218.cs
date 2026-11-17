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

    public interface IHandler11430
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11430(string Label, int Payload);

    public sealed class Handler11430 : IHandler11430
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate11431(IEnumerable<int> source)
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
    public struct Point11432
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid11432
    {
        public static Point11432 Make(int left, int right)
        {
            return new Point11432 { X = left, Y = right };
        }
    }
}
