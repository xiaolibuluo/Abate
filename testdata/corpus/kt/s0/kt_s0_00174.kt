package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull1420(): T? = if (size > 1) this[1] else null

fun render1420(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step1421")
fun step1421(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
data class Pair1422<A, B>(val first: A, val second: B)

fun combine1422(left: Int, right: Int): Pair1422<Int, Int> {
    return Pair1422(left + 22, right + 49)
}
fun <T> List<T>.secondOrNull1423(): T? = if (size > 1) this[1] else null

fun render1423(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry1424 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1424 = "registry-36"
    }
}
class Counter1425 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
object Registry1426 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1426 = "registry-50"
    }
}
data class Pair1427<A, B>(val first: A, val second: B)

fun combine1427(left: Int, right: Int): Pair1427<Int, Int> {
    return Pair1427(left + 57, right + 43)
}
class Counter1428 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
