package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4340")
fun step4340(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
data class Pair4341<A, B>(val first: A, val second: B)

fun combine4341(left: Int, right: Int): Pair4341<Int, Int> {
    return Pair4341(left + 28, right + 69)
}
class Counter4342 {
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
fun <T> List<T>.secondOrNull4343(): T? = if (size > 1) this[1] else null

fun render4343(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4344 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4344 = "registry-49"
    }
}
fun <T> List<T>.secondOrNull4345(): T? = if (size > 1) this[1] else null

fun render4345(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4346<A, B>(val first: A, val second: B)

fun combine4346(left: Int, right: Int): Pair4346<Int, Int> {
    return Pair4346(left + 63, right + 63)
}
object Registry4347 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4347 = "registry-70"
    }
}
fun <T, R> List<T>.mapChecked4348(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
