package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20340 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
object Registry20341 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20341 = "registry-70"
    }
}
data class Pair20342<A, B>(val first: A, val second: B)

fun combine20342(left: Int, right: Int): Pair20342<Int, Int> {
    return Pair20342(left + 77, right + 56)
}
sealed class Outcome20343 {
    data class Success(val payload: String) : Outcome20343()
    data object Failure : Outcome20343()
}

fun describe20343(outcome: Outcome20343): String = when (outcome) {
    is Outcome20343.Success -> outcome.payload
    Outcome20343.Failure -> "failed"
}
class Counter20344 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull20345(): T? = if (size > 1) this[1] else null

fun render20345(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome20346 {
    data class Success(val payload: String) : Outcome20346()
    data object Failure : Outcome20346()
}

fun describe20346(outcome: Outcome20346): String = when (outcome) {
    is Outcome20346.Success -> outcome.payload
    Outcome20346.Failure -> "failed"
}
