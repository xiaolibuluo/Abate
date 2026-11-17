package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14590")
fun step14590(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
object Registry14591 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14591 = "registry-64"
    }
}
sealed class Outcome14592 {
    data class Success(val payload: String) : Outcome14592()
    data object Failure : Outcome14592()
}

fun describe14592(outcome: Outcome14592): String = when (outcome) {
    is Outcome14592.Success -> outcome.payload
    Outcome14592.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14593(): T? = if (size > 1) this[1] else null

fun render14593(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14594 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14594 = "registry-85"
    }
}
data class Pair14595<A, B>(val first: A, val second: B)

fun combine14595(left: Int, right: Int): Pair14595<Int, Int> {
    return Pair14595(left + 92, right + 29)
}
object Registry14596 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14596 = "registry-99"
    }
}
fun <T, R> List<T>.mapChecked14597(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry14598 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14598 = "registry-16"
    }
}
object Registry14599 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14599 = "registry-23"
    }
}
