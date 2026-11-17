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

    public async Task<int> Aggregate5220(IEnumerable<int> source)
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
    public struct Point5221
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 8;
    }

    public static class Grid5221
    {
        public static Point5221 Make(int left, int right)
        {
            return new Point5221 { X = left, Y = right };
        }
    }
    public interface IHandler5222
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5222(string Label, int Payload);

    public sealed class Handler5222 : IHandler5222
    {
        public void Handle(int value) { }
    }
    public struct Point5223
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 22;
    }

    public static class Grid5223
    {
        public static Point5223 Make(int left, int right)
        {
            return new Point5223 { X = left, Y = right };
        }
    }
}
