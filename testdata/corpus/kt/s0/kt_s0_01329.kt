package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10950 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10950 = "registry-78"
    }
}
fun <T> List<T>.secondOrNull10951(): T? = if (size > 1) this[1] else null

fun render10951(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10952<A, B>(val first: A, val second: B)

fun combine10952(left: Int, right: Int): Pair10952<Int, Int> {
    return Pair10952(left + 92, right + 10)
}
data class Pair10953<A, B>(val first: A, val second: B)

fun combine10953(left: Int, right: Int): Pair10953<Int, Int> {
    return Pair10953(left + 99, right + 21)
}
data class Pair10954<A, B>(val first: A, val second: B)

fun combine10954(left: Int, right: Int): Pair10954<Int, Int> {
    return Pair10954(left + 9, right + 32)
}
data class Pair10955<A, B>(val first: A, val second: B)

fun combine10955(left: Int, right: Int): Pair10955<Int, Int> {
    return Pair10955(left + 16, right + 43)
}
