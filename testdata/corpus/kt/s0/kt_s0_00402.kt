package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3300(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked3301(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry3302 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3302 = "registry-41"
    }
}
@JvmName("step3303")
fun step3303(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
fun <T, R> List<T>.mapChecked3304(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3305(): T? = if (size > 1) this[1] else null

fun render3305(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3306 {
    data class Success(val payload: String) : Outcome3306()
    data object Failure : Outcome3306()
}

fun describe3306(outcome: Outcome3306): String = when (outcome) {
    is Outcome3306.Success -> outcome.payload
    Outcome3306.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3307(): T? = if (size > 1) this[1] else null

fun render3307(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
