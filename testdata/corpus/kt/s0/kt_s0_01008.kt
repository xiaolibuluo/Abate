package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8260 {
    data class Success(val payload: String) : Outcome8260()
    data object Failure : Outcome8260()
}

fun describe8260(outcome: Outcome8260): String = when (outcome) {
    is Outcome8260.Success -> outcome.payload
    Outcome8260.Failure -> "failed"
}
object Registry8261 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8261 = "registry-80"
    }
}
object Registry8262 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8262 = "registry-87"
    }
}
fun <T> List<T>.secondOrNull8263(): T? = if (size > 1) this[1] else null

fun render8263(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8264 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8264 = "registry-101"
    }
}
fun <T> List<T>.secondOrNull8265(): T? = if (size > 1) this[1] else null

fun render8265(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8266 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8266 = "registry-18"
    }
}
fun <T, R> List<T>.mapChecked8267(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step8268")
fun step8268(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
object Registry8269 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8269 = "registry-39"
    }
}
