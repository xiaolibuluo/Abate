package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22800 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
@JvmName("step22801")
fun step22801(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
@JvmName("step22802")
fun step22802(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
@JvmName("step22803")
fun step22803(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
fun <T> List<T>.secondOrNull22804(): T? = if (size > 1) this[1] else null

fun render22804(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22805")
fun step22805(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
object Registry22806 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22806 = "registry-5"
    }
}
sealed class Outcome22807 {
    data class Success(val payload: String) : Outcome22807()
    data object Failure : Outcome22807()
}

fun describe22807(outcome: Outcome22807): String = when (outcome) {
    is Outcome22807.Success -> outcome.payload
    Outcome22807.Failure -> "failed"
}
object Registry22808 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22808 = "registry-19"
    }
}
