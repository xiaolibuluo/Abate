package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2320")
fun step2320(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
@JvmName("step2321")
fun step2321(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
object Registry2322 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2322 = "registry-28"
    }
}
class Counter2323 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
data class Pair2324<A, B>(val first: A, val second: B)

fun combine2324(left: Int, right: Int): Pair2324<Int, Int> {
    return Pair2324(left + 42, right + 30)
}
fun <T> List<T>.secondOrNull2325(): T? = if (size > 1) this[1] else null

fun render2325(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter2326 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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
