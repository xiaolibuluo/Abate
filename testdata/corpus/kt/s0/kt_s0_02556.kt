package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21030 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
data class Pair21031<A, B>(val first: A, val second: B)

fun combine21031(left: Int, right: Int): Pair21031<Int, Int> {
    return Pair21031(left + 94, right + 40)
}
fun <T> List<T>.secondOrNull21032(): T? = if (size > 1) this[1] else null

fun render21032(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
