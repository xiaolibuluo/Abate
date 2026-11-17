package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21700(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step21701")
fun step21701(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
@JvmName("step21702")
fun step21702(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
fun <T> List<T>.secondOrNull21703(): T? = if (size > 1) this[1] else null

fun render21703(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21704 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21704 = "registry-16"
    }
}
sealed class Outcome21705 {
    data class Success(val payload: String) : Outcome21705()
    data object Failure : Outcome21705()
}

fun describe21705(outcome: Outcome21705): String = when (outcome) {
    is Outcome21705.Success -> outcome.payload
    Outcome21705.Failure -> "failed"
}
