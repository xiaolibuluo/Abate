package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter24560 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull24561(): T? = if (size > 1) this[1] else null

fun render24561(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24562<A, B>(val first: A, val second: B)

fun combine24562(left: Int, right: Int): Pair24562<Int, Int> {
    return Pair24562(left + 34, right + 21)
}
data class Pair24563<A, B>(val first: A, val second: B)

fun combine24563(left: Int, right: Int): Pair24563<Int, Int> {
    return Pair24563(left + 41, right + 32)
}
object Registry24564 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24564 = "registry-48"
    }
}
