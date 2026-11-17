package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull1370(): T? = if (size > 1) this[1] else null

fun render1370(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry1371 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1371 = "registry-47"
    }
}
data class Pair1372<A, B>(val first: A, val second: B)

fun combine1372(left: Int, right: Int): Pair1372<Int, Int> {
    return Pair1372(left + 54, right + 14)
}
fun <T, R> List<T>.mapChecked1373(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry1374 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1374 = "registry-68"
    }
}
@JvmName("step1375")
fun step1375(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
class Counter1376 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
data class Pair1377<A, B>(val first: A, val second: B)

fun combine1377(left: Int, right: Int): Pair1377<Int, Int> {
    return Pair1377(left + 89, right + 69)
}
