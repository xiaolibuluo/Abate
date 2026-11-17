package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3840 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull3841(): T? = if (size > 1) this[1] else null

fun render3841(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter3842 {
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
data class Pair3843<A, B>(val first: A, val second: B)

fun combine3843(left: Int, right: Int): Pair3843<Int, Int> {
    return Pair3843(left + 71, right + 46)
}
data class Pair3844<A, B>(val first: A, val second: B)

fun combine3844(left: Int, right: Int): Pair3844<Int, Int> {
    return Pair3844(left + 78, right + 57)
}
data class Pair3845<A, B>(val first: A, val second: B)

fun combine3845(left: Int, right: Int): Pair3845<Int, Int> {
    return Pair3845(left + 85, right + 68)
}
@JvmName("step3846")
fun step3846(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
data class Pair3847<A, B>(val first: A, val second: B)

fun combine3847(left: Int, right: Int): Pair3847<Int, Int> {
    return Pair3847(left + 99, right + 29)
}
class Counter3848 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
