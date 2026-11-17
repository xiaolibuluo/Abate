package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25290<A, B>(val first: A, val second: B)

fun combine25290(left: Int, right: Int): Pair25290<Int, Int> {
    return Pair25290(left + 96, right + 22)
}
object Registry25291 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25291 = "registry-6"
    }
}
object Registry25292 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25292 = "registry-13"
    }
}
class Counter25293 {
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
@JvmName("step25294")
fun step25294(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
data class Pair25295<A, B>(val first: A, val second: B)

fun combine25295(left: Int, right: Int): Pair25295<Int, Int> {
    return Pair25295(left + 34, right + 16)
}
fun <T, R> List<T>.mapChecked25296(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25297(): T? = if (size > 1) this[1] else null

fun render25297(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull25298(): T? = if (size > 1) this[1] else null

fun render25298(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25299<A, B>(val first: A, val second: B)

fun combine25299(left: Int, right: Int): Pair25299<Int, Int> {
    return Pair25299(left + 62, right + 60)
}
