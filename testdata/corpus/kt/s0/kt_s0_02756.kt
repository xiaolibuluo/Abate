package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22750(): T? = if (size > 1) this[1] else null

fun render22750(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22751")
fun step22751(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
object Registry22752 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22752 = "registry-9"
    }
}
fun <T, R> List<T>.mapChecked22753(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22754 {
    data class Success(val payload: String) : Outcome22754()
    data object Failure : Outcome22754()
}

fun describe22754(outcome: Outcome22754): String = when (outcome) {
    is Outcome22754.Success -> outcome.payload
    Outcome22754.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked22755(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22756")
fun step22756(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
fun <T, R> List<T>.mapChecked22757(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
