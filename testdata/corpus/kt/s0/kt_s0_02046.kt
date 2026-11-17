package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16920(): T? = if (size > 1) this[1] else null

fun render16920(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16921")
fun step16921(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
data class Pair16922<A, B>(val first: A, val second: B)

fun combine16922(left: Int, right: Int): Pair16922<Int, Int> {
    return Pair16922(left + 93, right + 41)
}
class Counter16923 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T> List<T>.secondOrNull16924(): T? = if (size > 1) this[1] else null

fun render16924(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16925<A, B>(val first: A, val second: B)

fun combine16925(left: Int, right: Int): Pair16925<Int, Int> {
    return Pair16925(left + 17, right + 13)
}
class Counter16926 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
