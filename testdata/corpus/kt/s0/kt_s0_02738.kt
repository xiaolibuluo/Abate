package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22600(): T? = if (size > 1) this[1] else null

fun render22600(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22601 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22601 = "registry-98"
    }
}
data class Pair22602<A, B>(val first: A, val second: B)

fun combine22602(left: Int, right: Int): Pair22602<Int, Int> {
    return Pair22602(left + 8, right + 52)
}
data class Pair22603<A, B>(val first: A, val second: B)

fun combine22603(left: Int, right: Int): Pair22603<Int, Int> {
    return Pair22603(left + 15, right + 63)
}
class Counter22604 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
